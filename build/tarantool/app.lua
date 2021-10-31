box.cfg{listen = 3302}
box.schema.user.passwd('pass')

box.once('init', function()
    s = box.schema.space.create('sessions')
    s:format({
        {name = 'sessionID', type = 'string'},
        {name = 'email', type = 'string'}
    })
    s:create_index('primary', {type = 'tree', parts = {'sessionID'}})

    k = box.schema.space.create('keys')
    k:format({
        {name = 'email', type = 'string'},
        {name = 'salt', type = 'string'}
    })
    k:create_index('primary', {type = 'tree', parts = {'email'}})
end)