package com.example.docker_demo.controller;
import com.example.docker_demo.dto.Movie;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.Arrays;
import java.util.List;

@RestController
@RequestMapping()
public class DockerDemoController {

    @GetMapping("/greet/{username}")
    public String greet (@PathVariable String username )
    {
        return "Hello " + username + ", Have a Great Day!!";
    }

    @GetMapping("movies")
    public List <Movie> myFavMov () {
        Movie[] ListOfMovies = {
                new Movie("Black Friday", 2004, "Drama", 8.5, "A gripping tale of the aftermath of the Mumbai Blast"),
                new Movie("Gupi Gayen Bagha Bayen", 1968, "Comedy", 8.9, "A hilarious tale of two friends and their adventures"),
                new Movie("Kiki's Delivery Service", 1989, "Animation", 8.0, "A heartwarming story of a young witch and her cat"),
                new Movie("Good Bye Lenin", 2003, "Drama", 7.8, "A touching story of a son trying to protect his mother from the truth"),
                new Movie("The Shawshank Redemption", 1994, "Drama", 9.3, "A story of hope and friendship in a prison"),
                new Movie("Django Unchained", 2012, "Western", 8.4, "A story of a freed slave on a mission to rescue his wife"),
        };
        return Arrays.asList(ListOfMovies);
    }
}
