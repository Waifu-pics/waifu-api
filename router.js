const config = require('./config')

module.exports = ({ db, app, s3 }) => {

    // * API
    
    // User endpoints
    require('./api/upload')({ db, app, config, s3 })

    // Admin endpoints
    require('./api/admin/auth/login')({ db, app, config, s3 })
    require('./api/admin/auth/verify')({ db, app, config, s3 })
    require('./api/admin/verify')({ db, app, config, s3 })
    require('./api/admin/listfile')({ db, app, config, s3 })

    // For every endpoint in config, create an endpoint
    config.endpoints.map(endpoint => {
        app.get(`/api/img/${endpoint}`, async (req, res) => {
            let randomFile = await db.collection("uploads").aggregate([{ $match: { "type": endpoint, "verified": true } }, { $sample: { size: 1 } } ]).toArray()
            res.json({
                'url': config.url + randomFile[0].file
            })      
        })
    })

    // * FRONTEND

    let gridPoint = JSON.parse(JSON.stringify(config.endpoints)).splice(1, config.endpoints.length)
    console.log(gridPoint);

    // Grid Frontend for SFW
    app.get('/', async(req, rep) => {
        // let collectionSize = await db.collection("uploads").countDocuments({"type": "sfw", "verified": true})
        rep.renderMin('grid', { data: await db.collection("uploads").aggregate([{ $match: {"type": "sfw", "verified": true} }, { $sample: { size: 100 } }]).toArray(), config: config })
    })

    config.endpoints.map(endpoint => {
        app.get(`/${endpoint}`, async(req, rep) => {
            rep.renderMin('grid', { data: await db.collection("uploads").aggregate([{ $match: {"type": endpoint, "verified": true} }, { $sample: { size: 100 } }]).toArray(), config: config })
        })
    })

    // User Frontend

    app.get('/docs', async(req, rep) => {
        rep.renderMin('docs', { endpoints : config.endpoints })
    })

    app.get('/upload', async(req, rep) => {
        rep.renderMin('upload', { maxUploadSize : config.maxUploadSize, endpoints : config.endpoints })
    })

    app.get('/admin/login', async(req, rep) => {
        rep.renderMin('admin/login')
    })

    app.get('/admin', async(req, rep) => {
        rep.redirect('/admin/login')
    })

    app.get('/admin/dash', async(req, rep) => {
        rep.renderMin('admin/dash')
    })

    // 404 Page
    app.get('*', function(req, rep){ 
        rep.renderMin('404'); 
    }) 
}