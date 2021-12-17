import {BrowserRouter, Routes, Route} from "react-router-dom";

import './App.css';
import PageNotFound from "./Components/404";

import Footer from './Components/Footer';
import Home from "./Components/Home";
import Movie from "./Components/Movie";
import Navbar from './Components/Navbar';


function App() {
  return (
  <div className='App'>
    <Navbar/>
    <BrowserRouter>
      <Routes>
        <Route path = "/" element = {<Home/>}/>
        <Route path = "/movie/:id" element = {<Movie/>}/>
        <Route path = "*" element = {<PageNotFound/>}/>
      </Routes>
    </BrowserRouter>
    <Footer/>
  </div>
  );
}

export default App;
