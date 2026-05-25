<script lang="ts">
  import LinkIcon from "@/icons/link-icon.svelte";
  import { css, cx } from "@emotion/css";
  import { card, flexbox, fontSize, spacing, typography } from "@/styled";
  import colors from "@/styled/colors";

  import image from "@/assets/avatar.png";
  import Avatar from "@/components/avatar.svelte";
  import Input from "@/components/input.svelte";
</script>

<img
  src="https://dy8n3onijof8f.cloudfront.net/static/images/wp/03740b1e7e77ad332349.svg"
  alt=""
/>

<div class={cx(spacing.pxMD, flexbox.flex, flexbox.flexCol, flexbox.gapSM)}>
  <div class={card.wrapper}>
    <div class={cx(card.content)}>
      <div
        class={cx(
          flexbox.flex,
          typography.fontBold,
          flexbox.alignItemsCenter,
          flexbox.gapXS,
        )}
      >
        <LinkIcon />
        <span
          class={cx(
            typography.textSm,
            typography.fontBold,
            typography.uppercase,
          )}>Link Halaman Kamu</span
        >
      </div>
      <a
        class={cx(
          typography.underline,
          typography.fontSemiBold,
          colors.primary,
          fontSize.textXs,
        )}
        href="https://sobuzz.id/dadanhdyt">https://sobuzz.id/dadanhdyt</a
      >
    </div>
  </div>

  <!-- website -->
  <div class={card.wrapper}>
    <div
      class={cx(
        css({
          borderBottomWidth: 2,
          borderBottomStyle: "solid",
          borderBottomColor: "var(--color-disabled)",
        }),
        spacing.pSM,
      )}
    >
      <h1 class={cx(typography.fontBold, typography.textMd)}>PROFILE</h1>
    </div>

    <div class={cx(card.content)}>
      <div class={cx(flexbox.flex, flexbox.gapMD, flexbox.alignItemsCenter)}>
        <Avatar />
        <div
          class={cx(
            css({ flex: 1 }),
            flexbox.flex,
            flexbox.flexCol,
            flexbox.gapSM,
          )}
        >
          <Input size="sm" placeholder="Nama Kreator" />
          <Input size="sm" placeholder="@username" />
        </div>
      </div>
    </div>
  </div>
</div>
