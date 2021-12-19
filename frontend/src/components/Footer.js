export default function Footer() {
    return (
        
    <div class = "bg-dark text-white p-1 pl-3 pr-3 mt-4">
    <div class = "row p-3">
        <div class = "col-md-6 container pl-md-5 pr-md-5 pt-4">
            <h3>
                About Code Scarlet
            </h3>
            <p class = "ml-2 mr-2">
                If you want to add any movie or series you like which is not in the list, smash that contribute button up there. Cheers!
            </p>
            <p class = "ml-2 mr-2">
                Find an issue with this page? <a href = "https://github.com/shrn01/code-scarlet">Raise issue</a>
            </p>
        </div>
        <div class = "col-md-4 container pr-md-5 pl-md-5 pt-4">
            <h3>
                Quick Links
            </h3>
            <p class = "footer-links">
                <ul class="list-unstyled pl-2">
                    <li>
                        <a class="text-white" href = "{{ url_for('index') }}">Home</a>
                        
                    </li>
                    <li>
                        <a class="text-white" href = "{{ url_for('random') }}">Random Suggestion</a>
                        
                    </li>
                    <li>
                        <a class = "text-white" href = "{{ url_for('genres') }}">All Genres</a>

                    </li>
                    <li>
                        <a class="text-white" href = "{{ url_for('contribute') }}">Contribute</a>
                        
                    </li>
                    <li>
                        <a class="text-white" href = "{{ url_for('contributors') }}">Contributors</a>
                        
                    </li>
                    <li>
                        <a class="text-white" href = "https://github.com/shrn01/code-scarlet">Github</a>
                    </li>
                  </ul>
                
            </p>
            
        </div>
    </div>
</div>
    );
}
