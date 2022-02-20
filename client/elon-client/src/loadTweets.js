import React, { useEffect, useState } from 'react';
import Highcharts, { setOptions } from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import axios from 'axios';


const LoadTweets = () => {
    const [options, setOptions] = useState({
        title: {
            text: "Tweets per day"
        },
        series: [{ data: null }],
    });
    useEffect(() => {
        axios.get("http://127.0.0.1:9090/api/tweetsperday")
            .then(data => {

                data.data.forEach(function (el) {
                    el._id = Date.parse(el._id)
                });
                const newdata = data.data.map(el => [el._id, el.count])
                setOptions({ series: [{ data: newdata }], xAxis: { type: 'datetime' } });

            });
    }, [])
    return <HighchartsReact highcharts={Highcharts} options={options} />
}

export default LoadTweets