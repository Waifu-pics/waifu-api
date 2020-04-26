const config = require('./config')

module.exports = ({ db, app, s3 }) => {

    // * API
    
    // User endpoints
    require('./api/upload')({ db, app, config, s3 })

    // * FRONTEND

    // User Frontend
    app.get('/', async(req, rep) => {
        rep.renderMin('index')
    })

    // User Frontend
    app.get('/upload', async(req, rep) => {
        rep.renderMin('upload', { maxUploadSize : config.maxUploadSize })
    })

    app.get('/nsfw', async(req, rep) => {
        let collectionSize = await db.collection("uploads").countDocuments({"type": "nsfw", "verified": true})
        rep.renderMin('nsfw', { data: await db.collection("uploads").aggregate([{ $match: {"type": "nsfw"} }, { $sample: { size: collectionSize } }]).toArray(), config: config })
    })

    app.get('/sfw', async(req, rep) => {
        let collectionSize = await db.collection("uploads").countDocuments({"type": "sfw", "verified": true})
        rep.renderMin('sfw', { data: await db.collection("uploads").aggregate([{ $match: {"type": "sfw"} }, { $sample: { size: collectionSize } }]).toArray(), config: config })
    })

    // // 404 Page
    // app.get('*', function(req, rep){ 
    //     rep.renderMin('404'); 
    // }) 
}