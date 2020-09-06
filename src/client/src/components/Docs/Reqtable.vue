<template>
  <v-simple-table>
    <thead>
      <tr>
        <th class="text-left">Field</th>
        <th class="text-left">Data Type</th>
        <th class="text-left">Description</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>{{field}}</td>
        <td>{{type}}</td>
        <td><slot></slot></td>
      </tr>
    </tbody>
  </v-simple-table>
</template>

<script>
export default {
  name: "reqtable",
  props: ["type", "field"],
}
</script>
