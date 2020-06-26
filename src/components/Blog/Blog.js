import PropTypes from "prop-types";
import React from "react";

import Item from "./Item";

const Blog = props => {
  const { posts, theme } = props;

  return (
    <React.Fragment>
      <main className="main">
        <h2>ブログ一覧</h2>
        <ul className="grid">
          {posts.map(post => {
            const {
              node,
              node: {
                fields: { slug }
              }
            } = post;
            return <Item key={slug} post={node} theme={theme} />;
          })}
        </ul>
      </main>

      {/* --- STYLES --- */}
      <style jsx>{`
        .main {
          padding: 0 ${theme.space.inset.default};
          h2 {
            text-align: center;
            font-size: 2.1em;
            margin-top: 50px;
          }
        }

        ul {
          list-style: none;
          margin: 0 auto;
          padding: ${`calc(${theme.space.default} * 1.5) 0 calc(${theme.space.default} * 0.5)`};
        }

        @above tablet {
          .main {
            padding: 0 ${`0 calc(${theme.space.default} * 1.5)`};
          }
          ul {
            max-width: ${theme.text.maxWidth.tablet};
          }
        }
        @above desktop {
          .main {
            .grid {
              display: grid;
              grid-template-columns: 33% 33% 33%;
            }
          }
          ul {
            max-width: ${theme.text.maxWidth.desktop};
          }
        }
      `}</style>
    </React.Fragment>
  );
};

Blog.propTypes = {
  posts: PropTypes.array.isRequired,
  theme: PropTypes.object.isRequired
};

export default Blog;
