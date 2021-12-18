export default function Navbar() {
    return (
        <div>
            <div className = "container-fluid bg-dark text-white">
                <nav className = "navbar">
                    <div className = 'text-white'>
                        <h3 className="mb-1 mt-1">Code Scarlet</h3>
                    </div>
                    <div className = 'd-none d-md-flex btn-group float-right'>
                        <div className = 'btn btn-secondary'><a href = "{{ url_for('index') }}" className = 'text-white'>Home</a></div>
                        <div className = 'btn btn-secondary'><a href = "{{ url_for('random') }}" className = 'text-white'>Random movie</a></div>
                        <div className = 'btn btn-secondary'><a href = "{{ url_for('genres') }}" className = 'text-white'>Genres</a></div>
                        <div className = 'btn btn-secondary'><a href = "{{url_for('contribute') }}" className = 'text-white'>Contribute</a></div>
                    </div>
                    {/* <div className = 'd-md-none btn-group float-right'>
                        <button type="button" className="btn btn-secondary dropdown-toggle shadow-none" data-toggle="dropdown">Menu
                            <span className="caret"></span>
                        </button>
                        <ul className = 'dropdown-menu dropdown-menu-right shadow border-0' role = 'menu'>
                            <li><a href = "{{ url_for('index') }}" className = 'text-center dropdown-item' role="menuitem">Home</a></li>
                            <li><a href = "{{ url_for('random') }}" className = 'text-center dropdown-item' role="menuitem">Random movie</a></li>
                            <li><a href = "{{url_for('genres') }}" className = 'text-center dropdown-item' role="menuitem">Genres</a></li>
                            <li><a href = "{{url_for('contribute') }}" className = 'text-center dropdown-item' role="menuitem">Contribute</a></li>
                        </ul>
                    </div> */}
                </nav>
            </div>
        </div>
    )
}
