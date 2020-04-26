const config = require('./config')

module.exports = ({ db, app, s3 }) => {

    // * API
    
    // User endpoints
    require('./api/upload')({ db, app, config, s3 })
    require('./api/admin/auth/login')({ db, app, config, s3 })
    require('./api/admin/auth/verify')({ db, app, config, s3 })
    require('./api/admin/verify')({ db, app, config, s3 })
    require('./api/admin/listfile')({ db, app, config, s3 })

    // * FRONTEND

    // User Frontend
    app.get('/', async(req, rep) => {
        rep.renderMin('index')
    })

    // User Frontend
    app.get('/upload', async(req, rep) => {
        rep.renderMin('upload', { maxUploadSize : config.maxUploadSize })
    })

    app.get('/admin/login', async(req, rep) => {
        rep.renderMin('admin/login')
    })

    app.get('/admin/dash', async(req, rep) => {
        rep.renderMin('admin/dash')
    })

    app.get('/nsfw', async(req, rep) => {
        let collectionSize = await db.collection("uploads").countDocuments({"type": "nsfw", "verified": true})
        rep.renderMin('grid', { data: await db.collection("uploads").aggregate([{ $match: {"type": "nsfw", "verified": true} }, { $sample: { size: collectionSize } }]).toArray(), config: config })
    })

    app.get('/sfw', async(req, rep) => {
        let collectionSize = await db.collection("uploads").countDocuments({"type": "sfw", "verified": true})
        rep.renderMin('grid', { data: await db.collection("uploads").aggregate([{ $match: {"type": "sfw", "verified": true} }, { $sample: { size: collectionSize } }]).toArray(), config: config })
    })

    // // 404 Page
    // app.get('*', function(req, rep){ 
    //     rep.renderMin('404'); 
    // }) 
}