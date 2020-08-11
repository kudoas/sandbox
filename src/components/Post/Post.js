import React from "react";
import PropTypes from "prop-types";
import Img from "gatsby-image";
import "prismjs/themes/prism-okaidia.css";

import asyncComponent from "../AsyncComponent";
import Headline from "../Article/Headline";
import Bodytext from "../Article/Bodytext";
import Meta from "./Meta";
import Comments from "./Comments";
import NextPrev from "./NextPrev";

const Share = asyncComponent(() =>
  import("./Share")
    .then(module => {
      return module.default;
    })
    .catch(error => {
      console.log("asyncComponentError: ", error);
    })
);

const Post = props => {
  const {
    post,
    post: {
      html,
      fields: { prefix, slug },
      frontmatter: {
        title,
        author,
        category,
        createdAt,
        cover: {
          children: [{ fluid }]
        }
      }
    },
    facebook,
    next: nextPost,
    prev: prevPost,
    theme
  } = props;

  return (
    <React.Fragment>
      <header>
        <div className="timer">{createdAt}</div>
        <Headline title={title} theme={theme} />
        <Meta prefix={createdAt} author={author} category={category} theme={theme} />
        <div className="gatsby-image-outer-wrapper">
          <Img fluid={fluid} />
        </div>
      </header>
      <Bodytext html={html} theme={theme} />
      <footer>
        <Share post={post} theme={theme} />
        <NextPrev next={nextPost} prev={prevPost} theme={theme} />
        <Comments slug={slug} facebook={facebook} theme={theme} />
      </footer>
      <style jsx>{`
        .timer {
          margin: 10px;
          font-size: 1.3em;
          color: #666;
          letter-spacing: 0.5px;
          text-align: center;
        }
        .gatsby-image-outer-wrapper {
          margin: 10px 15px;
          border: solid 1px #eee;
          picture {
            img {
              border-radius: 10px;
            }
          }
        }
      `}</style>
    </React.Fragment>
  );
};

Post.propTypes = {
  post: PropTypes.object.isRequired,
  authornote: PropTypes.string.isRequired,
  facebook: PropTypes.object.isRequired,
  next: PropTypes.object,
  prev: PropTypes.object,
  theme: PropTypes.object.isRequired
};

export default Post;
