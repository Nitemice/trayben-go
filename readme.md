# TrayBen

*TrayBen* is a basic system tray icon utility.

It is intended for use as a status indicator in various scenarios, with similar usage to `notify-send`.

It has no functionality, apart from the ability to display any given arguments as tooltip text.

## Example

One possible scenario is to show a tray icon when a given long-running task finishes.

```shell
slow-script.sh && trayben "script finished"

```
