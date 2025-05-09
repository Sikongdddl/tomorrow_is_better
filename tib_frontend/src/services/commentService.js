import axios from 'axios';

//const baseURL = 'http://localhost:8080';
const baseURL = 'http://58.196.157.31:23457';
const commentService = axios.create({
    baseURL,
    timeout: 10000,
})

export const addComment = (user_id, topic_id, content) => {
    return commentService.post('comment/add',{
        user_id: user_id,
        topic_id:topic_id,
        content:content
    });
}

export const getComment = (topic_id) => {
    return commentService.post('comment/list',{
        topic_id: topic_id
    });
}

export const delComment = (comment_id) => {
    return commentService.post('comment/delete',{
        comment_id: comment_id
    });
}