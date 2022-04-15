use core::fmt;
use std::fmt::Display;

use actix_web::{http::StatusCode, HttpResponse, ResponseError};
use serde::Serialize;

/// Standard basicmessage response
#[derive(Serialize, Debug)]
pub struct MessageResponse {
    #[serde(skip_serializing)]
    internal_code: StatusCode,

    code: u16,
    message: String,

    #[serde(skip_serializing_if = "Option::is_none")]
    error: Option<String>,
}

impl MessageResponse {
    /// Create new message response
    pub fn new(code: StatusCode, message: &str) -> Self {
        Self {
            internal_code: code,
            code: code.as_u16(),
            message: message.to_string(),
            error: None,
        }
    }

    /// New internal server error response
    pub fn internal_server_error(error: &str) -> Self {
        let mut response = MessageResponse::new(
            StatusCode::INTERNAL_SERVER_ERROR,
            "There was a problem processing your request",
        );

        response.error = Some(error.to_string());

        response
    }

    /// Create new unauthorized error response
    pub fn unauthorized_error() -> Self {
        MessageResponse::new(
            StatusCode::UNAUTHORIZED,
            "You are not authorized to make this request",
        )
    }

    /// Create new bad request error response
    pub fn bad_request() -> Self {
        MessageResponse::new(StatusCode::BAD_REQUEST, "You sent an invalid request")
    }

    /// Explicit convert to actix HttpResponse type
    pub fn http_response(&self) -> HttpResponse {
        HttpResponse::build(self.internal_code).json(self)
    }
}

/// Implicit From convert to actix HttpResponse type
impl From<MessageResponse> for HttpResponse {
    fn from(response: MessageResponse) -> Self {
        response.http_response()
    }
}

impl Display for MessageResponse {
    fn fmt(&self, fmt: &mut fmt::Formatter) -> fmt::Result {
        write!(fmt, "code: {}, message: {}", self.code, self.message)
    }
}

impl ResponseError for MessageResponse {
    fn status_code(&self) -> StatusCode {
        self.internal_code
    }

    fn error_response(&self) -> HttpResponse {
        self.http_response()
    }
}
