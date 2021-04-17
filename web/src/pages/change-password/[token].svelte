<script lang="ts">
  import Navbar from "../../components/Navbar.svelte"
  import {params, redirect} from "@roxi/routify"
  let newPassword: string;
  let error: Error;
  
  async function chagePassword() {
    try {
      const res = await fetch("http://localhost:1323/change-password", {
        method: "POST",
        body: JSON.stringify({
          token: $params["token"],
          newPassword: newPassword
        }),
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
            },
        credentials: "include",
      })
  
      if (res.ok) {
        $redirect("../home")
      } else {
        const err = await res.json();
        error = new Error(err.message)
      }
    } catch(err) {
      console.log(err)
      error = new Error("something wrong happend in the request")
    }
  }
</script>

<Navbar />
<main>
  <form action="#" on:submit|preventDefault={chagePassword}>
    <h1>Forgot Password</h1>
    <div class="formcontainer">
      <label for="email"><strong>New Password</strong></label>
      <input type="text" placeholder="enter password" bind:value={newPassword} />
      {#if error}
        <div class="error-container">
          <p style="color: #003049;">{error.message}</p>
        </div>
      {/if}
    </div>
    <button type="submit">Change Password</button>
  </form>
</main>

<style>
    main {
    margin-left: 30%;
    margin-right: 30%;
  }

  form {
    border: 5px solid #f1f1f1;
  }

  input[type="text"] {
    width: 100%;
    padding: 16px 8px;
    margin: 10px 0;
    display: inline-block;
    border: 1px solid #ccc;
    border-radius: 10px;
    box-sizing: border-box;
  }

  button {
    background-color: #003049;
    color: white;
    padding: 14px 0;
    margin: 10px 0;
    border: none;
    border-radius: 10px;
    width: 100%;
  }

  h1 {
    text-align: center;
    font-size: 18;
  }

  button:hover {
    opacity: 0.8;
  }

  .formcontainer {
    text-align: left;
    margin: 24px 50px 12px;
  }

  .error-container {
    background-color: #c1121f;
    /* padding: 1px 0px; */
  }
</style>

