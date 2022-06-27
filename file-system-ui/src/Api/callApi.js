import axios from 'axios';
import * as Config from './../Constants/Config';

axios.defaults.baseURL = `${Config.API_URL}`;

const CallApi = (endpoint, method = 'GET', data) => {
    return axios({
        method: method,
        url: `/${endpoint}`,
        data: data,
        headers: { 'content-type': 'application/x-www-form-urlencoded' }
    })
}

export default CallApi