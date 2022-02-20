import React from 'react';
import { render } from 'react-dom';
import { BrowserRouter as Router, Routes, Route, NavLink } from "react-router-dom";
import MostLeastLiked from './loadMostLeastLiked';
import LoadTweets from './loadTweets';
import LoadRetweets from './loadRetweets';
import TweetsHourly from './loadTweetsHourly';
import './index.css';

const App = () => {

  return <div><Router>
    <NavLink className="button" to="/loadtweets">Load Tweets Per Day</NavLink>
    <NavLink className="button" to="/loadretweets">Load Retweets Per Day</NavLink>
    <NavLink className="button" to="/mostleastliked">Load Most/Least Liked Tweet</NavLink>
    <NavLink className="button" to="/loadtweetshourly">Tweets Per Hour</NavLink>
    <Routes>
      <Route exact path="/mostleastliked" element={<MostLeastLiked/>} />
      <Route exact path="/loadtweets" element={<LoadTweets/>} />
      <Route exact path="/loadretweets" element={<LoadRetweets/>} />
      <Route exact path="/loadtweetshourly" element={<TweetsHourly/>} />
    </Routes>

  </Router>
  </div>

};

render(<App />, document.getElementById('root'));
