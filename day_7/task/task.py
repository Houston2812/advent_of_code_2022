import os


def read_input(folder_name: str, file_name: str) -> str:
    """Read in the input file

    Args:
        folder_name (str): Folder name containing the input file
        file_name (str): File name of the input file

    Returns:
        str: String of the datastream
    """

    # Set directories for folder and input file
    main_folder = os.path.join(os.getcwd(), folder_name)
    input_file = os.path.join(main_folder, file_name)

    # Read in input file as a list
    with open(input_file, 'r', encoding='utf-8') as f:
        f_content = [line.rstrip() for line in f.readlines()]

    return f_content


def change_directory(act_dir: str, command: str) -> str:
    """Change the active directory based on a "$ cd ..." command

    Args:
        act_dir (str): The current active directory
        command (str): The "$ cd ..." commmand

    Returns:
        str: The new current directory
    """

    # Identify the new directory to change to
    new_dir = command.split('$ cd ')[1]

    # If the command is to move out one level
    if new_dir == '..':
        # If the active directory is only one-level deep (ex: '/a)
        if act_dir.count('/') == 1:
            # Return the root directory
            return '/'
        # If the active directory is more than one-level deep
        else:
            # Remove the last sub-directory
            return '/'.join(act_dir.split('/')[:-1])
    # If the commmand is to move in one level
    else:
        # If there is no active directory or if the new directory is the root
        if act_dir == '' or new_dir == '/':
            # Return the root directory
            return '/'
        # If the active directory is the root directory
        elif act_dir == '/':
            # Add the sub-directory
            return act_dir + new_dir
        # If the active directory is more than one-level deep
        else:
            # Add the sub-directory into the active directory
            return act_dir + '/' + new_dir


def create_system_directory(outputs: list) -> dict:
    """Create a dictionary representation of all files in the system

    Args:
        outputs (list): A list of all commands and their outputs

    Returns:
        dict: A nested dictionary representation of the system directory
    """

    sys_dir = {}
    act_dir = ''
    for output in outputs:
        # For "$ cd" commands
        if output[:4] == '$ cd':
            # Update the current active directory
            act_dir = change_directory(act_dir, output)
        # For "$ ls" commands
        if output[:4] == '$ ls':
            # Initialize the nested dictionary for the directory
            sys_dir[act_dir] = {'children': [], 'files': [], 'size': 0}
        # For the output lines for each "$ ls" command
        if output[0] != '$':
            # If the file is a directory
            if output[:3] == 'dir':
                # Add the sub-directory to the children of the active directory
                sys_dir[act_dir]['children'].append(output[4:])
            # If the file is a file
            else:
                # Separate the file size and file name
                f_size_str, f_name = output.split(' ')
                # Append a tuple of the name and file size to the file list
                sys_dir[act_dir]['files'].append((f_name, int(f_size_str)))
                # Add the file size to the directory size
                sys_dir[act_dir]['size'] += int(f_size_str)

    return sys_dir


def calc_total_size(sys_dir: dict) -> dict:
    """Calculate the total size of each directory, including the file sizes of
    all its children directories

    Args:
        sys_dir (dict): A nested dictionary representation of the system
            directory

    Returns:
        dict: Dictionary of all directories and their total sizes including
            all children directories
    """

    # Initialize new dictionary with 0 default value for every directory
    all_sizes = {dir_path: 0 for dir_path in sys_dir.keys()}

    # Loop through both dictionaries and add the file sizes for all children
    for dir_path in all_sizes.keys():
        for dir, dir_info in sys_dir.items():
            if dir_path in dir:
                all_sizes[dir_path] += dir_info['size']

    return all_sizes


def calc_size_below_threshold(all_sizes: dict, threshold: int) -> int:
    """Calculate the total fize size of the directories below a certain
    threshold

    Args:
        all_sizes (dict): Dictionary of all directories and their total sizes
            including all children directories
        threshold (int): File size threshold value

    Returns:
        int: Total file size of the directories below a certain threshold
    """

    # Find the sum of file sizes below a threshold value
    below_threshold_sum = 0
    for size in all_sizes.values():
        if size <= threshold:
            below_threshold_sum += size

    return below_threshold_sum


def find_minimum_deletable_directory(all_sizes: dict,
                                     total_disk_space: int,
                                     update_size: int) -> tuple:
    """Find the smallest directory that can be deleted to clear up enough
    space for the update

    Args:
        all_sizes (dict): Dictionary of all directories and their total sizes
        total_disk_space (int): Total disk space available to the filesystem
        update_size (int): Size of the update file

    Returns:
        tuple: (Name of the minimum deletable directory, size of minimum
            deletable directory)
    """

    # Calculate required extra space to install system udpate
    used_space = total_disk_space - all_sizes['/']
    required_extra_space = update_size - used_space

    # Find all directories that can be deleted to free up enough space
    deletable_dirs = {dir_path: size for dir_path, size in all_sizes.items()
                      if size >= required_extra_space}

    # Find the smallest directory that can be deleted to free up enough space
    min_deletable_directory = min(deletable_dirs, key=deletable_dirs.get)

    return (min_deletable_directory, all_sizes[min_deletable_directory])


if __name__ == '__main__':

    outputs = read_input('.', 'input_task')
    sys_dir = create_system_directory(outputs)
    all_sizes = calc_total_size(sys_dir)

    # Solution to part 1
    print(f'The sum of the total size of the directories below 100000 is:'
          f' {calc_size_below_threshold(all_sizes, 100000)}.')

    min_deletable_directory = find_minimum_deletable_directory(all_sizes,
                                                               70000000,
                                                               30000000)

    # Solution to part 2
    print(f'The smallest directory that can be deleted is '
          f"{min_deletable_directory[0]}, and that directory's total size is "
          f'{min_deletable_directory[1]}.')