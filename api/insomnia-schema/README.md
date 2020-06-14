# Feel the Movies API v1

Get movies and tv shows recommendations!

I decided to open these endpoints to all developers who want to test
their web or mobile applications.

If you have any doubt, open an issue or send me an e-mail: xorycx@gmail.com.

## Endpoints

**Base URL:** https://feelthemovies.com.br/v1

**Image URL:** https://image.tmdb.org/t/p/

Supported sizes by TMDB:

- w45
- w92
- w154
- w300
- w780
- w1280
- h632
- original

Build the image URL like this:

**The image URL + Size + Image Key**

https://image.tmdb.org/t/p/w500/aMpyrCizvSdc0UIMblJ1srVgAEF.jpg

## Recommendations

This endpoint will provide the last 10 recommendations.

**GET - /recommendations?page=1**

If you want to paginate, increase the page number, like this:

**GET - /recommendations?page=2**

## Recomendation Items

This endpoint will provide all movies or tv shows of a specific recommendation.

**GET - /recommendation_items/80**

## Search Recommendations

This endpoint will provide 10 results based on your search. You can search using keywords
or genres, like this:

**GET - /search_recommendation?query=drama&page=1**

**GET - /search_recommendation?query=fight&page=1**

If you want to paginate, increase the page number, like this:

**GET - /search_recommendation?query=drama&page=2**

**GET - /search_recommendation?query=fight&page=2**

## Specific Recommendation 

This endpoint will return a single result of a given recommendation.

**GET - /recommendation/80**