package com.example.docker_demo.dto;


/*
    * @author Soumen Das
    * @date 2025-040-07
    * create class based on this
    *
	Title    string  `json:"title"`
	Year     int     `json:"year"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Synopsis string  `json:"synopsis"`
 */
public class Movie {
    private String title;
    private int year;
    private String genre;
    private double rating;
    private String synopsis;
    // Constructor
    public Movie(String title, int year, String genre, double rating, String synopsis) {
        this.title = title;
        this.year = year;
        this.genre = genre;
        this.rating = rating;
        this.synopsis = synopsis;
    }
    // Getters and Setters
    public String getTitle() {
        return title;
    }
    public void setTitle(String title) {
        this.title = title;
    }
    public int getYear() {
        return year;
    }
    public void setYear(int year) {
        this.year = year;
    }
    public String getGenre() {
        return genre;
    }
    public void setGenre(String genre) {
        this.genre = genre;
    }

    public double getRating() {
        return rating;
    }
    public void setRating(double rating) {
        this.rating = rating;
    }
    public String getSynopsis() {
        return synopsis;
    }
    public void setSynopsis(String synopsis) {
        this.synopsis = synopsis;
    }
}
