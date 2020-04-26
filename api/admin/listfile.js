const argon = require('argon2')

module.exports = ({ db, app, config }) => {
    app.post('/api/admin/listfile', async (req, res) =>{
        const { token } = req.body

        if (!token) {
            return res.status(400).send('Invalid credentials!')
        }

        const Admins = db.collection('admins')

        const tokenExists = Boolean(await Admins.findOne({ token }))

        const Uploads = db.collection('uploads')
        if (tokenExists) {
            const results = ( await Uploads.find({"verified": false}).sort({_id:-1}).toArray() ).map( file => { return {file: file.file, type: file.type} } )

            res.status(200).json(results)
        } else {
            return res.status(400).send('Invalid credentials!')
        }
    })
}