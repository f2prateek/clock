# clock

A clock abstraction providing access to the current time.
It's primary purpose is to allow alternate clocks to be plugged in for tests.

3 implementations are provided out of the box:

1. `Default`: Provides the current local time.
2. `Fixed`: Always provides a fixed time.
3. `Offset`: Provides the time from a different clock with the specified duration added.
