<script lang="ts">
  import DuolingoIcon from "@/icons/duolingo-icon.svelte";
  import IconHome from "@/icons/icon-home.svelte";
  import { css } from "@emotion/css";

  let { children } = $props();

  const layoutManage = {
    wrapper: css({
      display: "grid",
      gridTemplateColumns: "250px 1fr",
      minHeight: "100vh",
    }),
    sidebar: {
      menu: css({
        padding: 10,
      }),
      wrapper: css({
        borderRight: "2px solid",
        borderColor: "#e5e5e5",
      }),
      nav_link: {
        wrapper: css({
          boxShadow: "inset 0 -4px #0000000f",
          border: "1px solid #d2d4d8",
          padding: 10,
          borderRadius: 10,
          gap: 10,
          ":hover": {
            background: "#f7f7f7",
          },
          cursor: "pointer",
          color: "#777777",
          display: "flex",
          alignItems: "center",
        }),
        content: css({
          display: "flex",
          flexDirection: "column",
          span: {
            ":first-child": {
              textTransform: "uppercase",
              fontWeight: 800,
            },
            ":last-child": {
              fontSize: 12,
            },
          },
        }),
      },
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
        flex:1,
        display:'grid',
        gridTemplateColumns:"1fr"
    }),
  };
</script>

<div class={layoutManage.wrapper}>
  <aside class={layoutManage.sidebar.wrapper}>
    <div class={layoutManage.sidebar.brand}>
      <DuolingoIcon />
    </div>
    <div class={layoutManage.sidebar.menu}>
      <div class={layoutManage.sidebar.nav_link.wrapper}>
        <div>
          <IconHome />
        </div>
        <div class={layoutManage.sidebar.nav_link.content}>
          <span>Dadan Hidayat</span>
          <span>@dadanhdyt - creator</span>
        </div>
      </div>
    </div>
  </aside>
  <main class={layoutManage.main}>
    {@render children?.()}
  </main>
</div>
