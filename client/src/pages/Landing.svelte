<script lang="ts">
</script>

<main>
  <div class="landing-card">
    <div class="landing-card-item">
      <h1>Paraphrase like a pro with A.I.</h1>
      <p>Generate tailored paraphrases for your reports, essays</p>
      <button>Get Started </button>
    </div>
    <div class="landing-card-item">
      <div
        class="landing-card-image"
        style="background-image: url('typing-mac-book.jpg')"
      />
    </div>
  </div>
  <div class="landing-card">
    <div class="landing-card-item">
      <div
        class="landing-card-image"
        style="background-image: url('essay-writing.jpg')"
      />
    </div>
    <div class="landing-card-item">
      <h1>Write your papers worry free, directly from your browser</h1>
      <button>Get Started </button>
    </div>
  </div>
</main>

<style>
  main {
    width: 100%;
    text-align: center;
    margin: 0 auto;
  }

  .landing-card {
    height: 95vh;
    margin: 5px 15px;
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    align-items: center;
    padding: 20px;
  }

  .landing-card-item {
    width: 45%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 10px;
    align-items: center;
  }

  .landing-card-image {
    width: 100%;
    height: 85vh;
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center;
    border-radius: 15px;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
