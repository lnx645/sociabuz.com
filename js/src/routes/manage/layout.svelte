<script lang="ts">
  import NavLink from "@/components/sidebar/nav-item.svelte";
  import { mq } from "@/core/utils";
  import DuolingoIcon from "@/icons/duolingo-icon.svelte";
  import FeaturesIcon from "@/icons/features-icon.svelte";
  import IconHeart from "@/icons/icon-heart.svelte";
  import IconHome from "@/icons/icon-home.svelte";
  import IconMenu from "@/icons/icon-menu.svelte";
  import ListIcon from "@/icons/list-icon.svelte";
  import MoreIcon from "@/icons/more-icon.svelte";
  import StoreIcon from "@/icons/store-icon.svelte";
  import { css } from "@emotion/css";

  let { children } = $props();

  const layoutManage = {
    wrapper: css(
      mq({
        display: "grid",
        gridTemplateColumns: ["1fr", "256px 1fr"],
        minHeight: "100vh",
      }),
    ),
    sidebar: {
      menu: css({
        display: "flex",
        flexDirection: "column",
        gap: 10,
        padding: 10,
      }),
      wrapper: css(
        mq({
          borderRight: "2px solid",
          borderColor: "#e5e5e5",
          display: ["none", "block"],
        }),
      ),
      brand: css({
        paddingTop: 24,
        paddingInline: 10,
        paddingBottom: 5,
        fontWeight: "bold",
        svg: {
          width: "50%",
        },
      }),
    },
    main: css({
      flex: 1,
      display: "grid",
      gridTemplateColumns: "1fr",
    }),
  };
</script>

<div class={layoutManage.wrapper}>
  <aside class={layoutManage.sidebar.wrapper}>
    <div class={layoutManage.sidebar.brand}>
      <DuolingoIcon />
    </div>
    <div class={layoutManage.sidebar.menu}>
      <NavLink
        isActive={true}
        title="PROFILE"
        description="@dadanhdyt - creator"
        icon={IconHome}
      />
      <NavLink
        title="TRANSAKSI"
        description="Kelola transaksi"
        icon={ListIcon}
      />
      <NavLink
        title="FEATURES"
        description="Overlay,SondBoard,dll"
        icon={FeaturesIcon}
      />
      <NavLink title="MORE" icon={MoreIcon} />
    </div>
  </aside>
  <main class={layoutManage.main}>
    <nav
      class={css(
        mq({
          height: 56,
          borderBottom: "2px solid var(--color-disabled)",
          display: ["flex", "none"],
          alignItems: "center",
          paddingInline: 10,
        }),
      )}
    >
      <IconMenu />
    </nav>
    {@render children?.()}
  </main>
</div>
