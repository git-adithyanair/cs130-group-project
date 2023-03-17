# GoodGrocer

Adithya Nair (605399973), Yili Liu (205376049), Angela Hu (105401638), Bryan Luo (605303956), Roshni Rao (005394742)\
UCLA CS 130 Winter 2023

See the [Wiki](../../wiki) for the User Manual

---

## Installation Process

### Frontend

```
git clone git@github.com:git-adithyanair/cs130-group-project.git
cd cs130-group-project/frontend/GoodGrocer
npm install --force
```

To run the frontend through expo:

```
npm start
```

Then, follow the instructions in the terminal to open the app.\
If you have a simulator installed, type in 'i' to the terminal to open the app on an iOS simulator.\
Otherwise, download the [expo app](https://apps.apple.com/us/app/expo-go/id982107779) and scan to QR code in the simulator to run it.

\
_Note:_ The app was built and tested for the iOS platform so it may not perform the same on Android. It is recommended to run the frontend on an iOS operating system.

### Backend (optional)

The server is already deployed at http://api.good-grocer.click and deploys every time the backend folder in this repository is updated on the main branch so there is no need to manually deploy after running the frontend. However, if you want to run the server locally:

**Prerequisites:**\
Have Docker installed and running

**Then:**

```
cd cs130-group-project/backend
make devup
```

This will start the server at `http://localhost:8080`.

**To quit the server:**

```
^C
make devdown
```

Running this make command ensures that docker removes the image created during the make devup.

---

## Triggering Build

**GitHub Actions** is used for our CI/CD pipeline. As such, every commit pushed to a PR will trigger our testing workflow, and every PR that is merged with our main branch will trigger both our testing and build/deploy workflows.
