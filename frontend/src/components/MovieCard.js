
export default function MovieCard(props) {
    console.log("In movieCard")
    return (
        <div>
            <h2>{props.movie["movie"]}</h2>
        </div>
    )
}
