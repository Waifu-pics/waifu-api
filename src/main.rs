use actix_web::{
    error::JsonPayloadError, http::StatusCode, middleware::Logger, web::JsonConfig, App,
    HttpRequest, HttpServer,
};

use colored::*;
use config::Config;

use crate::responses::MessageResponse;

mod config;
mod responses;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    std::env::set_var("RUST_BACKTRACE", "1");
    std::env::set_var("RUST_LOG", "actix_web=info,sqlx=error,waifu_api=info");
    env_logger::init();

    let config = Config::new();

    let _sentry_guard = sentry::init((
        config.sentry_url,
        sentry::ClientOptions {
            release: sentry::release_name!(),
            session_mode: sentry::SessionMode::Request,
            auto_session_tracking: true,
            attach_stacktrace: true,
            ..Default::default()
        },
    ));

    log::info!(
        "Started waifu-api on port {}",
        config.port.to_string().yellow()
    );
    HttpServer::new(move || {
        App::new()
            .wrap(
                sentry_actix::Sentry::new()
                    .into_builder()
                    .capture_server_errors(false)
                    .finish(),
            )
            .wrap(Logger::default())
            .app_data(JsonConfig::default().error_handler(
                |err: JsonPayloadError, _: &HttpRequest| {
                    actix_web::Error::from(match err {
                        JsonPayloadError::Deserialize(json_err) => {
                            MessageResponse::new(StatusCode::BAD_REQUEST, &json_err.to_string())
                        }
                        _ => MessageResponse::bad_request(),
                    })
                },
            ))
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}
