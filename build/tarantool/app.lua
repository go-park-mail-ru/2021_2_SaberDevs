box.cfg{listen = 3302}
box.schema.user.passwd('pass')
queue = require 'queue'

-- queue.tube.articleLikes:drop()
-- queue.tube.articleComments:drop()
--
-- queue.create_tube('articleLikes', 'fifo', {temporary = true})
-- queue.create_tube('articleComments', 'fifo', {temporary = true})

box.once('init', function()
    s = box.schema.space.create('sessions')
    s:format({
        {name = 'sessionID', type = 'string'},
        {name = 'login', type = 'string'}
    })
    s:create_index('primary', {type = 'tree', parts = {'sessionID'}})

    k = box.schema.space.create('keys')
    k:format({
        {name = 'login', type = 'string'},
        {name = 'salt', type = 'string'}
    })
    k:create_index('primary', {type = 'tree', parts = {'login'}})

    sub = box.schema.space.create('subscriptions')
    s:format({
        {name = 'login', type = 'string'},
        {name = 'endpoint', type = 'string'},
        {name = 'auth', type = 'string'},
        {name = 'p256dh', type = 'string'}
    })
    sub:create_index('primary', {type = 'tree', parts = {'login'}})

    box.queue.create_tube('articleLikes', 'fifo', {temporary = true})
    box.queue.create_tube('articleComments', 'fifo', {temporary = true})

end)

function articleLikesPut(a)
    return queue.tube.articleLikes:put(a)
end

function articleLikesTake()
    r = queue.tube.articleLikes:take()
    if r ~= nil then
        queue.tube.articleLikes:ack(r.task_id)
    end
    return r
end

function articleCommentPut(a)
    return queue.tube.articleComments:put(a)
end

function articleCommentTake()
    r = queue.tube.articleComments:take()
    if r ~= nil then
        queue.tube.articleComments:ack(r.task_id)
    end
    return r
end


