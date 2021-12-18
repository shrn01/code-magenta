import {BrowserRouter, Routes, Route} from "react-router-dom";

import './App.css';
import PageNotFound from "./components/404";

import Footer from './components/Footer';
import Home from "./components/Home";
import Movie from "./components/Movie";
import Navbar from './components/Navbar';


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
