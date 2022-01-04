import React, { Component } from 'react'
import axios from 'axios'


let Context = React.createContext()


export class ContextProvider extends Component {
    state = {
        movieList: [],
        movieDict: {},
        watchList: []
    }

    render() {
        return (
            <Context.Provider value={this.state}>
                {this.props.children}
            </Context.Provider>
        )
    }

    componentDidMount() {
        axios.get("http://localhost:8000/api/movies").then(
            response => {
                this.setState({
                    movieList: response.data["movie_list"]
                })
            }
        )
    }
}


export let ContextConsumer = Context.Consumer