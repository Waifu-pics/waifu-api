const fileUpload = require('express-fileupload')
const cryptoRandomString = require('crypto-random-string')
const path = require('path')
const fs = require('fs')

module.exports = ({ db, app, config, s3 }) => {
    app.use(fileUpload({
        limits: { fileSize: config.maxUploadSize * 1024 * 1024 },
        abortOnLimit: true,
        createParentPath: true
    }))

    app.post('/api/upload', async (req, res) => {
        const { type } = req.headers
        if (!type) {
            return res.status(400).send("Invalid type provided!")
        } else if(['sfw', 'nsfw'].includes(type)) {            
            const Uploads = db.collection('uploads')
        if (req.files.uploadFile == null || Object.keys(req.files.uploadFile).length === 0) {
            return res.status(400).send('File not uploaded!')
        } else {
            let uploadFile = req.files.uploadFile
            let md5 = uploadFile.md5

            if (Boolean(await Uploads.findOne({ md5 })) == true) {
                return res.status(400).send('File already exists!')
            } else {
                let randomString
                let file

                const extension = path.extname(uploadFile.name)
                if (['png', 'jpg', 'jpeg', 'gif'].includes(extension.replace('.', ""))) {
                    
                    randomString = cryptoRandomString({length: config.fileLength, type: 'url-safe'})
                    file = (randomString + extension)

                    // Reroll filename if exists
                    while (Boolean(await Uploads.findOne({ file })) || randomString.includes (".")) {
                        randomString = cryptoRandomString({length: config.fileLength, type: 'url-safe'})
                        file = (randomString + extension)
                    }

                // Upload file to server and send response
                uploadFile.mv(config.uploadDir + file).then(async function() {
                    // S3 upload 
                    const params = {
                        Bucket: config.s3.bucket,
                        Key: file,
                        Body: fs.readFileSync(config.uploadDir + file),
                        ACL: 'public-read',
                        ContentType: uploadFile.mimetype
                    }
                    s3.upload(params, async function(s3Err) {
                        if (s3Err) throw s3Err

                        await Uploads.insertOne({ file, md5, type, "verified": false })
                        fs.unlinkSync(config.uploadDir + file)
                        return res.json({
                            'url': config.url + file
                        })                                   
                    })
                })

                } else {
                    return res.status(400).send('Invalid image!')
                }
            }
        }
        } else {
            return res.status(400).send("Invalid type provided!")
        }

    })
}