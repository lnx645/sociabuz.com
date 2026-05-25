import { mq } from "@/core/utils";
import { css } from "@emotion/css";

const wrapper = css(
  mq({
    borderRadius: 10,
    overflow: "hidden",
    maxWidth: 520,
    border: "2px solid var(--color-disabled)",
    color: "var(--color-text)",
  }),
);

const content = css(mq({ fontSize: 14, padding: 10, fontWeight: 600 }));

const card = {
  wrapper,
  content,
};
export  {card};
