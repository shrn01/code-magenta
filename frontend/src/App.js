import {BrowserRouter, Routes, Route} from "react-router-dom";

import './App.css';
import PageNotFound from "./components/404";

import Footer from './components/Footer';
import Home from "./components/Home";
import Movie from "./components/Movie";
import Navbar from './components/Navbar';

import "./styles/bootstrap.min.css"


function App() {
  return (
  <div className='App'>
    
    <BrowserRouter>
      <Navbar/>
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
