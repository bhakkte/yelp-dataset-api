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
