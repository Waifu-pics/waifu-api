let isAdmin = async function (req, res, next, db) {
    const Admins = db.collection('admins')
        let token = req.headers.token
        if(token && typeof token === "string" && Boolean(await Admins.findOne({token}))) {
        return true
    }        
    return false


    next()
}