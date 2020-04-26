const argon = require('argon2')

module.exports = ({ db, app, config }) => {
    app.post('/api/admin/auth/login', async (req, res) =>{
        const { username, password } = req.body

        if (!username || !password) {
            return res.status(400).send('The username/password you entered is incorrect!')
        }

        const Admins = db.collection('admins')
        if (Boolean(await Admins.findOne({ username }))) {
            const user = await Admins.findOne({ username })
            if (await argon.verify(user.password, password)) {
                const { token } = await Admins.findOne({ username })
                res.status(200).send(token)
            } else {
                res.status(400).send('The username/password you entered is incorrect!')
            }
        } else {
            res.status(400).send('The username/password you entered is incorrect!')
        }
    })
}