import {BrowserRouter, Routes, Route} from "react-router-dom";

import './App.css';
import PageNotFound from "./components/404";
import { ContextProvider } from "./components/Context";
import Contribute from "./components/Contribute";

import Footer from './components/Footer';
import Home from "./components/Home";
import Movie from "./components/Movie";
import Navbar from './components/Navbar';
import Watchlist from "./components/Watchlist";

import "./styles/bootstrap.min.css"

/**
 * Add Footer later when all other stuff is done
 */
function App() {
  return (
  <ContextProvider>
  <div className='App'>
    
    <BrowserRouter>
      <Navbar/>
      <Routes>
        <Route path = "/" element = {<Home/>}/>
        <Route path = "/movie/:id" element = {<Movie/>}/>
        <Route path = "/contribute" element = {<Contribute/>}/>
        <Route path = "/watchlist" element = {<Watchlist/>}/>
        <Route path = "*" element = {<PageNotFound/>}/>
      </Routes>
    </BrowserRouter>
  </div>
  </ContextProvider>
  );
}

export default App;
