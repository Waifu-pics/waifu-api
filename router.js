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

    // Frontend endpoint
    config.endpoints.map(endpoint => {
        app.post(`/api/${endpoint}`, async(req, res) => {
            const body = req.body
            let arrEx = []
            if (body.exclude) {
                arrEx = body.exclude
            }

            let data = await db.collection("uploads").aggregate([{ $match: { "type": endpoint, "verified": true, file: { $nin: arrEx } } }, { $sample: { size: 30 } }]).toArray()

            data = data.map((image) => {
                return image.file
            })

            res.json({
                'data': data
            })
        })
    })

    // * FRONTEND

    // Grid Frontend for SFW
    app.get('/', async(req, rep) => {
        // let collectionSize = await db.collection("uploads").countDocuments({"type": "sfw", "verified": true})
        const endpoint = "sfw"
        rep.renderMin('grid', { config: config, endpoint: endpoint })
    })

    config.endpoints.map(endpoint => {
        app.get(`/${endpoint}`, async(req, rep) => {
            rep.renderMin('grid', { config: config, endpoint: endpoint })
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