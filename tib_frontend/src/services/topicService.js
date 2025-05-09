import axios from 'axios';

//const baseURL = 'http://localhost:8080';
const baseURL = 'http://58.196.157.31:23457';
const topicService = axios.create({
    baseURL,
    timeout: 10000,
})

export const fetchTopics = () => {
    return topicService.get('getAllTopics');
}

export const getTopicById = (topic_id) => {
    return topicService.post('getTopicById',{
        topic_id: topic_id
    });
}

export const createTopic = (topic) =>
{
    return topicService.post('createTopic', {
        title: topic.title,
        event_time: topic.event_time,
        location: topic.location,
        creator_id: topic.creator_id
    });
}

export const joinTopic = (user_id, topic_id) =>
{
    return topicService.post('joinTopic', {
        user_id: user_id,
        topic_id: topic_id // 虽然在路径中已经传了，但后端你的逻辑要求也在 JSON 体中传一次
    });
}

export const leaveTopic = (user_id, topic_id) => {
    return topicService.post('leaveTopic',{
        user_id: user_id,
        topic_id: topic_id
    });
}