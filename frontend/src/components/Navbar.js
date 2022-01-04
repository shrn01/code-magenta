import {Link} from "react-router-dom"


export default function Navbar() {
    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light pl-2 pr-2">
            <Link className="navbar-brand" to="/">Code Scarlet</Link>
            <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarNavDropdown">
                <ul className="navbar-nav">
                <li className="nav-item active">
                    <Link className="nav-link" to="/">Home</Link>
                </li>
                <li className="nav-item">
                    <Link className="nav-link" to="/contribute">Contribute</Link>
                </li>
                <li className="nav-item">
                    <Link className="nav-link" to="/watchlist">Watchlist</Link>
                </li>
                <li className="nav-item">
                    <Link className="nav-link" to="/genres">Genres</Link>
                </li>
                </ul>
                
            </div>
            <form className="d-flex mr-2 ml-2">
                <input className="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search"/>
                <button className="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
            </form>
        </nav>
    )
}
