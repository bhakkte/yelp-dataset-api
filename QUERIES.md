# Queries

The following is a small collection of few interesting mongodb queries you can
run against the dataset.

Cites in the dataset and their review count

```javascript
db.businesses.group({
    initial: {
        count: 0
    },
    reduce: function(curr, res) {
        res.count++;
    },
    key: {
        city: 1
    },
    cond: {}
})
```

Businesses in Montreal (by city name)

```javascript
db.businesses.find({
    city: {$in: ["Montr\u00e9al", "Montreal"]}
}, {_id: 1, name: 1, city: 1, stars: 1, review_count: 1})
.sort({review_count: -1})
```

Businesses in a 8km sphere around Montreal (corner University and René-Lévesque)

```javascript
db.businesses.find({
    loc: {
        $near: {
            $geometry: {
                type: "Point" ,
                coordinates: [ -73.567256, 45.501689 ]
            },
            $maxDistance: 8000,
            $minDistance: 0
        }
    }
}, {_id: 1, name: 1, city: 1, stars: 1, review_count: 1})
.sort({review_count: -1})
```

Businesses in Edinburgh (by city name) grouped by star rating

```javascript
db.businesses.group({
    initial: {
        count: 0
    },
    reduce: function(curr, res) {
        res.count++;
    },
    key: {
        stars: 1
    },
    cond: {city: Edinburgh}
})
```

Queary in Mongodb finding star ratings for given cities and specific buisenesses. Database address: yack-mongodb.kiasaki.com:27017

Example: Nightlife, Casinos, Entertainment in Las Vegas

```javascript
db.businesses.group({
    initial: {
        count: 0
    },
    reduce: function(curr, res) {
        res.count++;
    },
    key: {
        stars: 1
    },
    cond: {
        city: {
            $in: ["Las Vegas", "North las Vegas", "South Las Vegas", "N Las Vegas", "North Las Vegas ", "N E Las Vegas", "N W Las Vegas", "Las Vegas ", "LAS VEGAS", "N. Las Vegas", "Las Vegas East", "C. Las Vegas"]
        },
        categories: {
            $elemMatch: {
                $in: ["Nightlife", "Casinos", "Casino", "Arts & Entertainment", "Cannabis Clinic", "Venues & Event Spaces", "Festivals", "Hotels"]
            }
        }
    }
})

```


Quick and dirty star count js script

```javascript
var rawstars = {
    "0" : {
        "stars" : 4,
        "count" : 2885
    },
    "1" : {
        "stars" : 4.5,
        "count" : 2028
    },
    "2" : {
        "stars" : 3,
        "count" : 1887
    },
    "3" : {
        "stars" : 2.5,
        "count" : 1201
    },
    "4" : {
        "stars" : 5,
        "count" : 1739
    },
    "5" : {
        "stars" : 3.5,
        "count" : 2928
    },
    "6" : {
        "stars" : 1.5,
        "count" : 263
    },
    "7" : {
        "stars" : 2,
        "count" : 551
    },
    "8" : {
        "stars" : 1,
        "count" : 147
    }
  };
var result = [];
for (key in rawstars) {
  var v = rawstars[key];
  result[v.stars * 2] = "("+v.stars+","+v.count+")";
}
console.log("[" + result.join(", ").slice(4) + "]")
```
