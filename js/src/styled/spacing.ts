import { css } from "@emotion/css";


const pSM = css({ padding: "8px" });
const pMD = css({ padding: "16px" });
const pLG = css({ padding: "24px" });

const pxMD = css({ paddingLeft: "16px", paddingRight: "16px" });
const pyMD = css({ paddingTop: "16px", paddingBottom: "16px" });


const mSM = css({ margin: "8px" });
const mMD = css({ margin: "16px" });
const mLG = css({ margin: "24px" });

const mtSM = css({ marginTop: "8px" });
const mbSM = css({ marginBottom: "8px" });

const spacing = {
  pSM, pMD, pLG,
  pxMD, pyMD,
  mSM, mMD, mLG,
  mtSM, mbSM
};

export default spacing;