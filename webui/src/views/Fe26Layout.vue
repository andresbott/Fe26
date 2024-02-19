<template>
  <div class="fe26-container">
    <header>
      <slot name="header"></slot>
    </header>
    <div class="left"></div>
    <div class="menu">
      <slot name="menu"></slot>
    </div>


    <div class="main">
      <slot></slot>
    </div>
    <div class="right"></div>

    <footer>
      <slot name="footer"></slot>
    </footer>
  </div>
</template>

<style>
.fe26-container {
  display: grid;
  //grid-template-areas:
  //  "header header header header"
  //  "left menu content right"
  //  "footer footer footer footer";
  //grid-template-rows: auto 1fr auto;
  //grid-template-columns: 1fr 2fr 3fr 1fr ;
  min-height: 100vh;
}

.fe26-container header {
  grid-area: header;
}

.fe26-container .left {
  grid-area: left;
  background: #478173;
}

.fe26-container .menu {
  grid-area: menu;
  background: #815959;
}

.fe26-container .main {
  grid-area: content;
  background: #6b8159;
}

.fe26-container .right {
  grid-area: right;
  background: #173d34;
}

.fe26-container footer {
  grid-area: footer;
}

/* Media queries for different screen sizes*/
@media (max-width: 600px) {
  .fe26-container {
    grid-template-areas:
      "header"
      "content"
      "footer";
    grid-template-rows: auto 1fr auto;
    grid-template-columns: 1fr;
  }
  .fe26-container .menu  {
    display:none;
  }
}


@media (min-width: 601px) {
  .fe26-container {
    grid-template-areas:
    "header header "
    "menu content "
    "footer footer";
    grid-template-rows: auto 1fr auto;
    grid-template-columns: 1fr 2fr ;
  }
  .fe26-container .menu,.fe26-container .main  {
    width: 100%
  }
}
/* add class full-with to not add borders at sizes >900px*/
@media (min-width: 901px) {
  .fe26-container {
    grid-template-areas:
    "header header header header"
          "left menu content right"
          "footer footer footer footer";
    grid-template-rows: auto 1fr auto;
    grid-template-columns: 1fr auto auto 1fr;
  }

  .fe26-container.full-with {
    grid-template-areas:
    "header header "
    "menu content "
    "footer footer";
    grid-template-rows: auto 1fr auto;
    grid-template-columns: 1fr 2fr ;
  }

  .fe26-container .menu {
    max-width: 400px;
    min-width:250px;
    margin: 0 0 0 auto;
  }

  .fe26-container .main {
    max-width: 900px;
    min-width: 650px;
    margin: 0 auto 0 0;
  }
  .fe26-container.full-with .menu,.fe26-container.full-with .main  {
    width: 100%
  }
}


</style>
<script setup>
</script>
