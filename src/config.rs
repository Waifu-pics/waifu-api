use dotenv::dotenv;
use rusoto_core::Region;
use std::{env, fmt::Debug, str::FromStr};

pub struct Config {
    pub port: u16,
    pub database_url: String,
    pub storage_url: String,
    pub sentry_url: String,
    pub s3_config: S3Config,
}

pub struct S3Config {
    pub bucket: String,
    pub access_key: String,
    pub secret_key: String,
    pub region: Region,
}

impl Config {
    pub fn new() -> Self {
        dotenv().ok();
        Config {
            port: get_env("PORT"),
            database_url: get_env("DATABASE_URL"),
            storage_url: get_env("STORAGE_URL"),
            sentry_url: get_env("SENTRY_URL"),
            s3_config: S3Config {
                bucket: get_env("S3_BUCKET"),
                access_key: get_env("S3_ACCESS_KEY"),
                secret_key: get_env("S3_SECRET_KEY"),
                region: Region::Custom {
                    name: get_env("S3_REGION"),
                    endpoint: get_env("S3_ENDPOINT"),
                },
            },
        }
    }
}

fn get_env<T>(var: &str) -> T
where
    T: FromStr,
    <T as FromStr>::Err: Debug,
{
    env::var(var)
        .expect(&format!("Missing environment variable {}", var))
        .parse::<T>()
        .expect(&format!(
            "Unable to parse {} as {}",
            var,
            std::any::type_name::<T>()
        ))
}
