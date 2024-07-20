<template>
  <v-container class="fill-height">
    <v-responsive class="align-center mx-auto" max-width="900">
      <div class="text-center">
        <h1 class="text-h2 font-weight-bold">🐙.☕️</h1>
        <h2 class="text-h5 font-weight-light">Password Checker</h2>
        <br />
      </div>

      <div class="py-4"></div>
      <v-row>
        <v-col cols="12">
          <v-card class="py-4" color="surface-variant" rounded="lg" variant="outlined">
            <v-form @submit.prevent="hashPassword">
              <h2 class="text-h5 font-weight-bold px-4">Enter a password</h2>
              <v-text-field v-model="password" label="Password" outlined type="password" class="px-4" required />
              <v-btn color="primary" class="mt-3 mx-4" type="submit">
                Check Password
              </v-btn>
            </v-form>
            <v-overlay opacity=".12" scrim="primary" :value="true" />
          </v-card>
        </v-col>
      </v-row>
      <!-- List of files -->
      <v-row>
        <v-col cols="12">
          <v-list>
            <v-subheader>Dictionaries containing the Password:</v-subheader>
            <v-list-item-group v-for="file in files" :key="file">
              <v-list-item>
                <v-list-item-content>
                  <v-list-item-title>{{ file }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-col>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { sha3_512 } from "js-sha3";

const password = ref("");
const files = ref([]);

async function hashPassword() {
  const hashedPassword = sha3_512(password.value);
  console.log(hashedPassword);
  fetch("https://api.checker.pops.cafe?hash=" + hashedPassword)
    .then((response) => response.json())
    .then((data) => {
      files.value = data.data.split(', ').sort()
        .map((file: string) => file.trim());
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}
</script>


<style scoped>
/* Responsive text styling */
.text-caption {
  padding: 0 20px;
  /* Adds horizontal padding for better text alignment */
  font-size: 1rem;
  /* Default font size */
}

/* Media queries for responsiveness */
@media (min-width: 600px) {
  .text-caption {
    font-size: 1.1rem;
    /* Slightly larger font size for tablets and small desktops */
  }
}

@media (min-width: 960px) {
  .text-caption {
    font-size: 1.25rem;
    /* Larger font size for larger screens */
  }
}
</style>
