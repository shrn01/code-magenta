import MovieCard from "./MovieCard";

export default function Home() {
    let movies = [
        "John Wick",
        "A Beautiful Place",
        "Avatar",
        "Baby Driver",
        "BirdBox",
        "Cast Away",
        "Deadpool",
        "Dont Breathe"
    ]
    console.log("In Home")
    return (
        <div>
            {
                movies.map(
                    movie => <MovieCard movie = {movie} key = {movies.indexOf(movie)}/>)
            }
        </div>
    );
}
