import { ContextConsumer } from "./Context";
import MovieCard from "./MovieCard";

export default function Home() {
    return (

        <>
        <ContextConsumer>
        {
            value => {
                return value.movieList.map(movie => <MovieCard movie = {movie} key = {movie["movie"]}></MovieCard>)
            }
        }
        </ContextConsumer>
</>
    );
}
