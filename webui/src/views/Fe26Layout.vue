<template>
  <div class="fe26-container">
    <header>
      <slot name="header"></slot>
    </header>

    <menu>
      <slot name="menu"></slot>
    </menu>

    <main>
      <slot></slot>
    </main>

    <footer>
      <slot name="footer"></slot>
    </footer>
  </div>
</template>

<style>
.fe26-container {
  background: #e8e8e8;
  display: grid;
  grid-template-areas:
    "header header"
    "sidebar fe26-main"
    "footer footer";
  grid-template-rows: auto 1fr auto;
  grid-template-columns: 1fr 3fr;
  min-height: 100vh;
}

.fe26-container header {
  background: #b7b7b7;
  grid-area: header;
}

.fe26-container menu {
  background: #1d062f;
  grid-area: sidebar;
  color: white;
}

.fe26-container main {
  background: #c2b8b0;
  grid-area: fe26-main;
}

.fe26-container footer {
  background: #9a1a67;
  grid-area: footer;
}

/* Media queries for different screen sizes */
@media (max-width: 600px) {
  .fe26-container {
    grid-template-areas:
      "header"
      "fe26-main"
      "sidebar"
      "footer";
    grid-template-rows: auto auto 1fr auto;
    grid-template-columns: 1fr;
  }
}

@media (min-width: 601px) {
  .fe26-container {
    grid-template-areas:
      "header header"
      "sidebar fe26-main"
      "footer footer";
    grid-template-rows: auto 1fr auto;
    grid-template-columns: auto auto;
  }
}

@media (min-width: 901px) {
  .fe26-container {
    grid-template-areas:
      "header header"
      "sidebar fe26-main"
      "footer footer";
    grid-template-rows: auto 1fr auto;
    grid-template-columns: auto auto;
  }

  .fe26-container menu {
    width: 300px;
    margin: 0 0 0 auto;
  }

  .fe26-container menu {
    width: 600px;
    margin: 0 auto 0 0;
  }
}
</style>
