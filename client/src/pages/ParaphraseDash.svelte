<script lang="ts">
  import { userStore } from "../store";
  import { getParaphrase, User } from "../infratructure/paraphrace";
  import Loader from "../components/Loader.svelte";

  let loading: boolean = false;
  let originalText: string = "";
  let resultText: string = "";
  let user: User;

  userStore.subscribe((v) => (user = v));

  async function parphraseIt() {
    if (originalText === "") {
      resultText = "Nothing to paraphrase here ¯\\_(ツ)_/¯ ";
      return;
    }
    loading = true;
    try {
      const res = await getParaphrase(user, originalText);
      resultText = res;
      loading = false;
    } catch (err) {
      loading = false;
      console.error(err);
      alert("Error getting paraphrase from API. Contact Support.");
    }
  }
</script>

<Loader show={loading} />
<main>
  <div>
    <h1>Source</h1>
    <textarea bind:value={originalText} />
  </div>
  <button on:click={parphraseIt}>Generate</button>
  <div>
    <h1>Results</h1>
    <textarea bind:value={resultText} readonly />
  </div>
</main>

<style>
  main {
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 20px;
  }

  textarea {
    width: 80vw;
    height: 25vh;
    min-width: 300px;
    min-height: 200px;
    background-color: #f0f0f0;
  }
</style>
