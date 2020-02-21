import React from "react";

const outerItem = props => {
  const { theme, item: { label, to, icon: Icon } = {}, onClick } = props;

  return (
    <React.Fragment>
      <li className={"hiddenItem" in props ? "hiddenItem" : "item"} key={label}>
        <a href={to} className={"hiddenItem" in props ? "inHiddenItem" : ""} onClick={onClick}>
          {Icon && <Icon />} {label}
        </a>
      </li>

      {/* --- STYLES --- */}
      <style jsx>{`
        .item,
        .showItem {
          background: transparent;
          transition: all ${theme.time.duration.default};
          display: flex;
          align-items: center;

          :global(a) {
            padding: ${theme.space.inset.s};
            display: flex;
            align-items: center;
          }

          :global(svg) {
            margin: 0 ${theme.space.inset.xs} 0 0;
            opacity: 0.3;
          }
        }

        :global(.itemList .hideItem) {
          display: none;
        }

        @from-width desktop {
          .item {
            :global(a) {
              color: ${theme.text.color.primary};
              padding: ${theme.space.inset.s};
              transition: all ${theme.time.duration.default};
              border-radius: ${theme.size.radius.small};
            }

            :global(.homepage):not(.fixed) & :global(a) {
              color: ${theme.color.neutral.white};
            }

            :global(a:hover) {
              color: ${theme.color.brand.primary};
              background: color(white alpha(-60%));
            }

            :global(svg) {
              transition: all ${theme.time.duration.default};
            }

            &:hover :global(svg) {
              fill: ${theme.color.brand.primary};
              opacity: 1;

              :global(.hero) & :global(svg) {
                fill: green;
              }
            }
          }

          .showItem {
            display: none;
          }

          .hiddenItem {
            text-align: left;
            padding: ${theme.space.xs};

            & :global(a.inHiddenItem) {
              color: ${theme.text.color.primary};
              &:hover {
                color: ${theme.color.brand.primary};
              }
            }
          }
        }
      `}</style>
    </React.Fragment>
  );
};

export default outerItem;
