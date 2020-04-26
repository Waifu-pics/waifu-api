const assert = require('assert')

module.exports = ({ db, app, config, s3 }) => {
    app.post('/api/admin/verify', async (req, res) =>{
        const { token, image, isVer } = req.body

        if (!token || !image || isVer === undefined) {
            return res.status(400).send('This data is invalid!')
        }

        const Admins = db.collection('admins')
        const Uploads = db.collection('uploads')

        const tokenExists = Boolean(await Admins.findOne({ token }))
        const uploadExists = Boolean(await Uploads.findOne({ "file": image }))

        
        if (tokenExists && uploadExists) {
            if (isVer) {
                let upload = Uploads.findOne({ "image": image })
                if (!upload.verified) {
                    Uploads.updateOne({ "file": image }, {$set: {"verified": true}})
                    res.status(200).send('Image has been verified!')
                } else {
                    res.status(400).send('Image already verified!')
                }
            } else {
                Uploads.deleteOne({ "file" : image })
                const params = { 
                    Bucket: config.s3.bucket, 
                    Key: image
                }
                s3.deleteObject(params, function(err) {
                    if (err) {
                        console.log(err)
                    }
                })
                res.status(200).send('Image Deleted!')
            }
        } else {
            res.status(400).send('Invalid!')
        }
    })
}