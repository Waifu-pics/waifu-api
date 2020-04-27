module.exports = {
    port: 8081,
    url: '',

    fileLength: 7, // Length of file names
    maxUploadSize: 20, // In MegaBytes
    uploadDir: '', // Directory to store temp uploads

    database: {
        mongoUrl: '',
        dbName: ''
    },

    s3: {
        bucket: '',
        accessKey: '',
        secretKey: '',
        endpoint: ''
    },

    endpoints: [
        "sfw",
        "nsfw"
    ]
}