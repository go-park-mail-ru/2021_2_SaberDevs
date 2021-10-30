box.cfg{listen = 3302}
box.schema.user.passwd('pass')

box.once('init', function()
    s = box.schema.space.create('sessions')
    s:format({
        {name = 'sessionID', type = 'string'},
        {name = 'email', type = 'string'}
    })
    s:create_index('primary', {type = 'tree', parts = {'sessionID'}})
end)