import React from "react";
import PropTypes from "prop-types";

const Bodytext = props => {
  const { html, theme } = props;

  return (
    <React.Fragment>
      <div className="bodytext" dangerouslySetInnerHTML={{ __html: html }} />

      <style jsx>{`
        .bodytext {
          animation-name: bodytextEntry;
          animation-duration: ${theme.time.duration.long};

          :global(body > *:first-child) {
            margin-top: 0 !important;
          }
          :global(body > *:last-child) {
            margin-bottom: 0 !important;
          }

          :global(h2),
          :global(h3) {
            margin: 1.5em 0 1em;
          }

          :global(h2) {
            line-height: ${theme.font.lineHeight.s};
            font-size: ${theme.font.size.l};
          }

          :global(h3) {
            font-size: ${theme.font.size.m};
            line-height: ${theme.font.lineHeight.m};
          }

          :global(p) {
            font-size: ${theme.font.size.s};
            line-height: ${theme.font.lineHeight.xxl};
            margin: 0 0 1.5em;
          }
          :global(ul) {
            list-style: circle;
            margin: 0 0 1.5em;
            padding: 0 0 0 1.5em;
          }
          :global(li) {
            margin: 0.7em 0;
            line-height: 1.5;
          }
          :global(a) {
            font-weight: ${theme.font.weight.bold};
            color: ${theme.color.brand.primary};
            text-decoration: underline;
          }
          :global(a.gatsby-resp-image-link) {
            border: 0;
            display: block;
            margin: 2.5em 0;
            border-radius: ${theme.size.radius.default};
            overflow: hidden;
            border: 1px solid ${theme.line.color};
          }
          :global(code.language-text) {
            background: ${theme.color.neutral.gray.c};
            text-shadow: none;
            color: inherit;
            padding: 0.1em 0.3em 0.2em;
            border-radius: 0.1em;
          }

          // table design
          :global(table) {
            border-spacing: 0;
            box-shadow: 0px 0px 5px 1px #eee;
            width: 100%;
            :global(tbody tr:nth-child(odd)) {
              background-color: #eee;
            }
            :global(th),
            :global(td) {
              text-align: center;
              width: 25%;
              padding: 15px 0;
            }
          }
          :global(table tr th :first-child),
          :global(table tr td :first-child) {
            marigin-top: 0;
          }
          :global(table tr th :last-child),
          :global(table tr td :last-child) {
            marigin-bottom: 0;
          }

          // blockquote
          :global(blockquote) {
            border-left: 4px solid #dddddd;
            padding: 0 15px;
            margin-bottom: 15px;
            color: #777777;
          }
          :global(blockquote > :first-child) {
            margin-top: 0;
          }
          :global(blockquote > :last-child) {
            margin-bottom: 0;
          }
          :global(img) {
            text-align: center;
            max-width: 70%;
            box-shadow: 0px 0px 5px 1px #eee;
            border-radius: 3px;
            display: block;
            margin: 20px auto;
          }

          :global(code, tt) {
            margin: 0 2px;
            padding: 0 5px;
            white-space: nowrap;
            border: 1px solid #eaeaea;
            background-color: #f8f8f8;
            border-radius: 3px;
          }
          :global(pre code) {
            margin: 0;
            padding: 0;
            white-space: pre;
            border: none;
            background: transparent;
          }
        }

        @keyframes bodytextEntry {
          from {
            opacity: 0;
          }
          to {
            opacity: 1;
          }
        }
      `}</style>
    </React.Fragment>
  );
};

Bodytext.propTypes = {
  html: PropTypes.string.isRequired,
  theme: PropTypes.object.isRequired
};

export default Bodytext;
