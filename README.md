# Express.js API Server with MongoDB Sample

## Abstruct

- Web API Server Sample
- Web Flamework : Express.js (https://expressjs.com/)
- ODM : Mongoose (https://mongoosejs.com/)

## Getting Started 

```
git pull git@github.com:Kudoas/express-api-sample.git
cd express-api-sample
```

Before setting up the server, you need to register with MongoDB and enter the URL in the `.env`file.

```
yarn run dev
```

## API Documents

### Post

#### Response example (Get All)

```
{
  "posts": [
    {
      "_id": "5f6583f39b1617658a2fd428",
      "title": "my post",
      "url": "https://www.google.com/",
      "text": "sample text",
      "createdAt": "2020-09-19T04:07:15.102Z",
      "updatedAt": "2020-09-19T04:07:15.102Z",
      "__v": 0
    }
  ]
}
```

