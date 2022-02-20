import React, { useEffect, useState } from 'react';
import { render } from 'react-dom';
import Highcharts, { setOptions } from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import axios from 'axios';
import { withRouter } from 'react-router';


const TweetsHourly = () => {
    const [options, setOptions] = useState({
        title: {
            text: "Tweets distribution per hours"
        },
        series: [{ data: null }],
    });

    useEffect(() => {
        axios.get("http://127.0.0.1:9090/api/tweetshourly")
            .then(data => {
                const newdata = data.data.map(el => [el._id, el.count])
                setOptions({
                    series: [{ data: newdata }], xAxis: {
                        categories: data.data.map(el => el._id)
                    }
                });
            });
        }, [])
        return <HighchartsReact highcharts={Highcharts} options={options} />
    }
    
    export default TweetsHourly
