module.exports = ({ db, app, config }) => {
    app.post('/api/admin/auth/verify', async (req, res) =>{
        const { token } = req.body

        if (!token) {
            res.status(400).send('This token is invalid!')
        }

        const Admins = db.collection('admins')

        const tokenExists = Boolean(await Admins.findOne({ token }))

        if (tokenExists) {
            res.status(200).send('This token is valid!')
        } else {
            res.status(400).send('This token is invalid!')
        }
    })
}