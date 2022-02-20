import React, { useEffect, useState } from 'react';
import { render } from 'react-dom';
import Highcharts, { setOptions } from 'highcharts';
import HighchartsReact from 'highcharts-react-official';
import axios from 'axios';
import { withRouter } from 'react-router';

const MostLeastLiked = () => {
  
    const [mostLiked, setMostLiked] = useState({});
    const [leastLiked, setLeastLiked] = useState({});
    const [showMostLeast, setShowMostLeast] = useState(false);
    useEffect( async () => {
      const result = await axios.get("http://127.0.0.1:9090/api/mostlikedtweet")
      setMostLiked(result.data[0])
      const result2 = await axios.get("http://127.0.0.1:9090/api/leastlikedtweet")
      setLeastLiked(result2.data[0])
      setShowMostLeast(true)
    });
    return <div>
    <h2>Most Liked Tweet</h2>
    <h5>Tweet:{mostLiked.tweet}</h5>
    <h5>Likes:{mostLiked.likescount}</h5>
    <h5>Date:{mostLiked.date}</h5>
    <h2>Least Liked Tweet</h2>
    <h5>Tweet:{leastLiked.tweet}</h5>
    <h5>Likes:{leastLiked.likescount}</h5>
    <h5>Date:{leastLiked.date}</h5>
  </div>
}
export default MostLeastLiked